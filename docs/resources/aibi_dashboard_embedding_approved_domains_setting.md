---
subcategory: "Settings"
---

# databricks_aibi_dashboard_embedding_approved_domains_setting Resource

-> This resource can only be used with a workspace-level provider!

The `databricks_aibi_dashboard_embedding_approved_domains_setting` resource allows you to specify the list of domains allowed for  [embedding of AI/BI Dashboards](https://learn.microsoft.com/en-us/azure/databricks/dashboards/admin/#manage-dashboard-embedding) into other sites.

## Example Usage

```hcl
resource "databricks_aibi_dashboard_embedding_access_policy_setting" "this" {
  aibi_dashboard_embedding_access_policy {
    access_policy_type = "ALLOW_APPROVED_DOMAINS"
  }
}

resource "databricks_aibi_dashboard_embedding_approved_domains_setting" "this" {
  aibi_dashboard_embedding_approved_domains {
    approved_domains = ["test.com"]
  }
  depends_on = [databricks_aibi_dashboard_embedding_access_policy_setting.this]
}
```

## Argument Reference

The resource supports the following arguments:

- `aibi_dashboard_embedding_approved_domains` block with following attributes:
  - `approved_domains` - (Required) the list of approved domains. To allow all subdomains for a given domain, use a wildcard symbol (`*`) before the domain name, i.e., `*.databricks.com` will allow to embed into any site under the `databricks.com`.

## Import

This resource can be imported by predefined name `global`:

```bash
terraform import databricks_aibi_dashboard_embedding_approved_domains_setting.this global
```

## Related Resources

The following resources are often used in the same context:

- [databricks_aibi_dashboard_embedding_access_policy_setting](databricks_aibi_dashboard_embedding_access_policy_setting.md) is used to control embedding policy.
