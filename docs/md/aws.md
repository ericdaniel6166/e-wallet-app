##### AWS
- AWS ECR
  - Amazon Elastic Container Registry
  - AWS managed container image registry service
  - automatically build using GitHub Action
- AWS IAM
  - AWS Identity and Access Management
  - AWS service, control access to AWS resources 
  - Identities
    - IAM User, is an entity that is created in AWS to represent the person or application 
    that used it to interact with AWS, consists of a name and credentials
    - IAM User Group, is a collection of users, help specify permissions for multiple users
    - IAM Role, is used to grant specific permissions to specific actors (
      - not have password, access key, belong to a group
      - can be associated with AWS resource
- AWS RDS
  - Amazon Relational Database Service
  - set up, operate, and scale a relational database in the AWS Cloud
- AWS Secrets Manager
  - secrets management service that helps protect access to applications, services, and IT resources 
  (ex: environment variables such as DB_SOURCE,TOKEN_SYMMETRIC_KEY, etc.)
  - easily rotate, manage, and retrieve database credentials, API keys, 
  and other secrets throughout their lifecycle
- AWS KMS 
  - AWS Key Management Service
  - a managed service help create and control the cryptographic keys that are used to protect data. (ex: AWS RDS)
- AWS CloudWatch
  - monitors AWS resources and the applications that run on AWS in real time (ex: AWS RDS)
- AWS EKS
  - Amazon Kubernetes Service
  - a managed service help use Kubernetes on AWS 
  without needing to install, operate, and maintain Kubernetes control plane or nodes
  - Kubernetes is an open source system 
  for automating the deployment, scaling, and management of containerized applications 
- AWS Route 53
  - a highly available and scalable DNS (Domain Name System) Service 
  - 3 main functions:
    - domain registration
    - DNS routing (Routing internet traffic to the resources for domain)
    - health checking (Check the health of resources, to verify that it's reachable, available, and functional)