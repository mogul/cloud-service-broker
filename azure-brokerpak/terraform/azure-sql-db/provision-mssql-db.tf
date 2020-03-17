variable server_name { type = string }
variable db_name { type = string }
variable labels { type = map }
variable pricing_tier { type = string }
variable cores { type = number }
variable storage_gb { type = number }
variable admin_username { type = string }
variable admin_password { type = string }
variable resource_group { type = string }

data "azurerm_sql_server" "azure_sql_db_server" {
  name                         = var.server_name
  resource_group_name          = var.resource_group
}

resource "azurerm_sql_database" "azure_sql_db" {
  name                = var.db_name
  resource_group_name = var.resource_group
  location            = data.azurerm_sql_server.azure_sql_db_server.location
  server_name         = data.azurerm_sql_server.azure_sql_db_server.name
  requested_service_objective_name = format("%s_Gen5_%d", var.pricing_tier, var.cores)
  max_size_bytes      = var.storage_gb * 1024 * 1024 * 1024
  tags = var.labels
}

output "sqldbName" {value = "${azurerm_sql_database.azure_sql_db.name}"}
output "sqlServerName" {value = "${data.azurerm_sql_server.azure_sql_db_server.name}"}
output "sqlServerFullyQualifiedDomainName" {value = "${data.azurerm_sql_server.azure_sql_db_server.fqdn}"}
output "databaseLogin" {value = "${var.admin_username}"}
output "databaseLoginPassword" {value = "${var.admin_password}"}
output "jdbcUrl" {
    value = format("jdbc:sqlserver://%s:1433;database=%s;user=%s;password=%s;Encrypt=true;TrustServerCertificate=false;HostNameInCertificate=*.database.windows.net;loginTimeout=30", 
                   data.azurerm_sql_server.azure_sql_db_server.fqdn,
                   azurerm_sql_database.azure_sql_db.name, 
                   var.admin_username, 
                   var.admin_password)
}
output "jdbcUrlForAuditingEnabled" {
    value = format("jdbc:sqlserver://%s:1433;database=%s;user=%s;password=%s;Encrypt=true;TrustServerCertificate=false;HostNameInCertificate=*.database.windows.net;loginTimeout=30", 
                   data.azurerm_sql_server.azure_sql_db_server.fqdn, 
                   azurerm_sql_database.azure_sql_db.name, 
                   var.admin_username, 
                   var.admin_password)
}
output "hostname" {value = "${data.azurerm_sql_server.azure_sql_db_server.name}"}
output "port" {value = 1433}
output "name" {value = "${azurerm_sql_database.azure_sql_db.name}"}
output "username" {value = "${var.admin_username}"}
output "password" {value = "${var.admin_password}"}
output "uri" {
    value = format("mssql://%s:1433/%s?encrypt=true&TrustServerCertificate=false&HostNameInCertificate=*.database.windows.net", 
                   data.azurerm_sql_server.azure_sql_db_server.fqdn, 
                   azurerm_sql_database.azure_sql_db.name)
}