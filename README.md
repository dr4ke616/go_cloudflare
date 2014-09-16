###Go CloudFlare
Something to be aware of, not all CloudFlare calls are currently supported. This is a new project.

## Example
Some examples on how to use Go Cloudflare. You can get your CloudFlare token in the Account section of your profile.

To get a single record using the record ID (you can find the ID on your cloudflare account):
```go
import (
    "log"
    "github.com/dr4ke616/go_cloudflare"
)

client, err := go_cloudflare.NewClient("<CLOUDFLARE_EMAIL>", "<CLOUDFLARE_TOKEN>")
if err != nil {
    log.Fatal("Problem with clouflare client: ", err)
}

record, err := client.RetrieveARecord("somedomain.com", "<ID_OF_RECORD>")
if err != nil {
    log.Fatal("Problem with clouflare client: ", err)
}

log.Println("Record ID: ", record.Id)
log.Println("Record Domain: ", record.Domain)
log.Println("Record Name: ", record.Name)
log.Println("Record Full Name: ", record.FullName)
log.Println("Record Value: ", record.Value)
log.Println("Record Types: ", record.Type)
log.Println("Record Priority: ", record.Priority)
log.Println("Record Ttl: ", record.Ttl)
```

## Current Status

**DNS Record Management**

| Function Call      | CloudFlare Call   | Description                                                   |
| ------------------ | ----------------- | --------------------------------------------------------------|
| CreateRecord       | rec_new           | Add a DNS record                                              |
| UpdateRecord       | rec_edit          | Edit a DNS record                                             |
| DestroyRecord      | rec_delete        | Delete a DNS record                                           |

**Access**

| Function Call      | CloudFlare Call   | Description                                                   |
| ------------------ | ----------------- | --------------------------------------------------------------|
| RetrieveAllRecords | rec_load_all      | Retrieve DNS Records of a given domain                        |
| Not Supported      | stats             | Retrieve domain statistics for a given time frame             |
| Not Supported      | zone_load_multi   | Retrieve the list of domains                                  |
| Not Supported      | zone_check        | Checks for active zones and returns their corresponding zids  |
| Not Supported      | ip_lkup           | Check threat score for a given IP                             |
| Not Supported      | zone_settings     | List all current setting values                               |

**Modify**

| Function Call      | CloudFlare Call   | Description                                                   |
| ------------------ | ----------------- | --------------------------------------------------------------|
| Not Supported      | sec_lvl           | Set the security level                                        |
| Not Supported      | cache_lvl         | Set the cache level                                           |
| Not Supported      | devmode           | Toggling Development Mode                                     |
| Not Supported      | fpurge_ts         | Clear CloudFlare's cache                                      |
| Not Supported      | zone_file_purge   | Purge a single file in CloudFlare's cache                     |
| Not Supported      | wl / ban / nul    | Whitelist/Blacklist/Unlist IPs                                |
| Not Supported      | ipv46             | Toggle IPv6 support                                           |
| Not Supported      | async             | Set Rocket Loader                                             |
| Not Supported      | minify            | Set Minification                                              |
| Not Supported      | mirage2           | Set Mirage2                                                   |
