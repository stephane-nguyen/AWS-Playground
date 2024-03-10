resource "aws_vpc" "tf_ec2_vpc" {
  cidr_block           = "10.123.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = {
    Name = "dev"
  }
}

resource "aws_subnet" "tf_ec2_vpc_public_subnet" {
  vpc_id                  = aws_vpc.tf_ec2_vpc.id
  cidr_block              = "10.123.1.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "us-east-1a"
  tags = {
    Name = "dev-public"
  }
}

resource "aws_internet_gateway" "tf_ec2_vpc_internet_gateway" {
  vpc_id = aws_vpc.tf_ec2_vpc.id
  tags = {
    Name = "dev-igw"
  }
}

resource "aws_route_table" "tf_ec2_vpc_public_rt" {
  vpc_id = aws_vpc.tf_ec2_vpc.id
  tags = {
    Name = "dev-public-rt"
  }
}

resource "aws_route" "default_route" {
  route_table_id         = aws_route_table.tf_ec2_vpc_public_rt.id
  destination_cidr_block = "0.0.0.0/0" // all ip addr will route to this internet gateway
  gateway_id             = aws_internet_gateway.tf_ec2_vpc_internet_gateway.id
}

resource "aws_route_table_association" "tf_ec2_vpc_public_assoc" {
  subnet_id      = aws_subnet.tf_ec2_vpc_public_subnet.id
  route_table_id = aws_route_table.tf_ec2_vpc_public_rt.id
}

resource "aws_security_group" "tf_ec2_sg" {
  name        = "dev_sg"
  description = "dev security group"
  vpc_id      = aws_vpc.tf_ec2_vpc.id
  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"          // allow all
    cidr_blocks = ["0.0.0.0/0"] // can set /32 to allow only one ip address
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_key_pair" "tf_ec2_auth" {
  key_name   = "tfdevenv"
  public_key = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIIljm2pYgHY122LTnn2c6zLjTfNvCR3i7Y+z8kF+0UEE GLOBAL+I566780@W-D9CZD63"
}

resource "aws_instance" "dev_node" {
  instance_type          = "t2.micro"
  ami                    = data.aws_ami.server_ami.id
  key_name               = aws_key_pair.tf_ec2_auth.key_name
  vpc_security_group_ids = [aws_security_group.tf_ec2_sg.id]
  subnet_id              = aws_subnet.tf_ec2_vpc_public_subnet.id
  user_data              = file("userdata.tpl")

  // Resize default size of the instance
  root_block_device {
    volume_size = 10 // GB, initially 8
  }
  tags = {
    Name = "dev-node"
  }
}