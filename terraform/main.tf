module "vpc" {
    source = "./vpc"
    private_cidr_block = var.private_cidr_block
    public_cidr_block = var.public_cidr_block
}

module "rds" {
    source = "./rds"
    private_cidr_block = var.private_cidr_block
    subnet_id = module.vpc.private_subnet_id
    vpc_id = module.vpc.vpc_id
}