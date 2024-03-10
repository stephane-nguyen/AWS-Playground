# Create a dev environment on AWS EC2 instance using Terraform for automation.

## Overview

This guide outlines the process of creating a development environment on an AWS EC2 instance using Terraform for automation. It covers the creation of various AWS components and provides useful Terraform commands.

### Terraform useful commands

- `terraform state list`: View list of resources created.
- `terraform show`: See all attributes of all resources.
- `terraform state show aws_vpc.mtc_vpc`: View all attributes for a specific resource.
- `terraform plan destroy`: Preview what will be destroyed.
- Add `-auto-approve` to skip confirmation.
- `terraform fmt`: Format the Terraform code.
- `terraform state list`: to see list of resources created
- `terraform console`: to go to the console to test variables (var."varName"), see IP addr EC2 instance (aws_instance.dev_node.public_ip)...
- `terraform apply -refresh-only`: to avoid re-creating all ressources when updating tf files...
- `terraform output`: to see outputs

## Explanation of AWS Resources and Their Purpose

1. **VPC (Virtual Private Cloud):**

   - **Purpose:** Provides an isolated virtual network environment in the cloud.
   - **Link:** The subnet and internet gateway are associated with the VPC.

2. **Subnet:**

   - **Purpose:** Divides the VPC's IP address range into smaller segments to organize resources.
   - **Link:** The subnet is associated with the VPC and specifies where EC2 instances will be launched.

3. **Internet Gateway:**

   - **Purpose:** Enables communication between instances in the VPC and the internet.
   - **Link:** The internet gateway is attached to the VPC to allow outbound and inbound internet traffic.

4. **Route Table:**

   - **Purpose:** Defines routing rules for traffic leaving the subnet.
   - **Link:** The route table is associated with the VPC and specifies the internet gateway as the target for outbound internet traffic.

5. **Route:**

   - **Purpose:** Specifies where network traffic should be directed based on destination IP addresses.
   - **Link:** Routes all IP traffic to the internet gateway, allowing instances in the subnet to communicate with the internet.

6. **Route Table Association:**

   - **Purpose:** Associates a subnet with a specific route table.
   - **Link:** The route table association ensures that the subnet uses the correct route table for routing traffic.

7. **Security Group:**

   - **Purpose:** Acts as a virtual firewall to control inbound and outbound traffic for instances.
   - **Link:** The security group is associated with the VPC and specifies rules for allowing traffic to and from the EC2 instances.

8. **Key Pair:**

   - **Purpose:** Provides secure access to EC2 instances by enabling SSH key-based authentication.
   - **Link:** The key pair is used to authenticate SSH connections to the EC2 instances.

9. **EC2 Instance:**
   - **Purpose:** Represents a virtual server in the cloud that runs applications and services.
   - **Link:** The EC2 instance is launched in the specified subnet and security group, using the key pair for SSH access. It communicates with the internet through the internet gateway and follows routing rules defined in the route table.

In summary, these resources work together to create a secure and functional environment for deploying EC2 instances in the AWS cloud. The VPC provides the network infrastructure, the subnet specifies where instances are located, the route table controls outbound traffic, the security group defines firewall rules, the key pair enables secure access, and the EC2 instance represents the virtual server itself.

### Steps for setups

**IAM User:**

- Create an IAM User with Admin rights.
- Download the AWS Extension on VS Code.
- Obtain the Access Keys for the IAM User by creating credentials profile in VS Code.
- On VSCode, CTRL + P, type ">Create credentials Profile" and mention the access keys of the IAM User

**AMI:**

- AMI : datasource query of the aws api to receive information needed to deploy a resource
  get AMI Owner.

  On AWS website, go to EC2 > Launch Instances > Check for a free tier Ubuntu and click on it. Copy the AMI ID (e.g ami-07d9b9ddc6cd8dd30)

  On the left tab > Image > AMI. Set "Owned by Me" by "Public Image" and enter in the search bar the AMI ID to get the Owner (e.g 099720109477)

**Key Pair:**

For windows: launch `ssh-keygen -t ed25519` in .ssh

A message will show up:
Generating public/private ed25519 key pair.
Enter file in which to save the key (C:\Users\username/.ssh/id_ed25519): C:\Users\username/.ssh/tfec2key
I did not bother with key passphrase and fingerprint so I just pressed enter

## Information

utilize userdata to bootstrap (= launch the script at the creation of the instance) to get an ec2 instance deployed with docker ready to go for all of our development needs.

.tpl = template

Get ec2 public ip address: terraform state show aws_instance.dev_node

`ssh -i ~/.ssl/keyname ubuntu@EC2PublicIpAdress`

Once connected, `docker --version` to see if it has been well installed
