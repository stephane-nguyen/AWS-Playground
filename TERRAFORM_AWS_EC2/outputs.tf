output "ec2_id" {
    value = aws_instance.dev_node.id
}

output "ec2_public_ip" {
    value = aws_instance.dev_node.public_ip
}

output "ec2_private_ip" {
    value = aws_instance.dev_node.private_ip
}

output "ec2_public_dns" {
    value = aws_instance.dev_node.public_dns
}

output "ec2_private_dns" {
    value = aws_instance.dev_node.private_dns
}

output "ec2_arn" {
    value = aws_instance.dev_node.arn
}

output "ec2_instance_type" {
    value = aws_instance.dev_node.instance_type
}