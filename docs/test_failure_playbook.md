# Terraform Google Conversion (TGC) Test Failure Playbook

This playbook helps diagnose and fix common issues encountered when running TGC tests. These tests often involve converting between CAI (Cloud Asset Inventory) JSON and HCL (Terraform).

---

## Debugging the Conversion Workflow

These tests check the accuracy of the conversions between Cloud Asset Inventory (CAI) format and Terraform HCL. The workflow consists of the following stages:

1.  **Initial Conversion (CAI to HCL):** The input CAI asset file (`Test.json`) is converted into a Terraform configuration file (`Test_export.tf`) using the `cai2hcl` tool.
    *   **Verify:** Does `Test_export.tf` correctly represent the resources defined in `Test.json`? Does the Terraform decoder or flattener cause issues?
2.  **Plan and Convert (HCL to CAI):** The generated `Test_export.tf` is used to create a Terraform plan, which is then converted into a CAI asset format file (`Test_roundtrip.json`) using the `tfplan2cai` tool.
    *   **Verify:** Does the terraform plan on `Test_export.tf` run without errors? Does `Test_roundtrip.json` match the original `Test.json`?

---

## Common Test Failures & Solutions

### 1. Conflicting Fields in HCL
*   **Symptom:** Error message like `Invalid combination of arguments ... next_hop_ilb: only one of ... can be specified`.
*   **Cause:** The CAI asset to HCL conversion (`cai2hcl`) produced HCL where multiple mutually exclusive fields are set.
*   **Solution:** Implement custom logic in a `tgc_decoder` to enforce exclusivity.
*   **Action:** Add `tgc_decoder: 'templates/tgc_next/decoders/xxxx.go.tmpl'` to the `custom_code` section of the resource's YAML file. Create the `.go.tmpl` file in `mmv1/templates/tgc_next/decoders` folder.

### 2. Missing Sensitive Fields
*   **Symptom:** Error message like `At least 1 'trust_anchors' blocks are required`.
*   **Cause:** Fields like `pemCertificate` don't exist in CAI assets as they are sensitive info, but are required in HCL.
*   **Solution:** Use a `tgc_decoder` to set missing sensitive fields to a placeholder like `"unknown"`.
*   **Example:** `TestAccIAMBetaWorkloadIdentityPoolProvider_x509` failed because `pemCertificate` was missing.

### 3. One-of Fields Not Set
*   **Symptom:** Error message: `'self_managed': one of managed, self_managed must be specified`.
*   **Cause:** Similar to conflicting fields, but where at least one must be present and the conversion failed to pick one.
*   **Solution:** Use a `tgc_decoder` to ensure the correct block is initialized.

### 4. Error Converting Round-Trip Config (API Calls)
*   **Symptom:** Error 404: `The resource ... was not found, 'notFound'`.
*   **Cause:** The HCL to CAI conversion (`custom_tgc_expand`) attempts to call the GCP API (to resolve instance names to IDs, etc.), but the resource doesn't exist in the test project yet.
*   **Solution:** Override the default expansion logic for the field causing the API call.
*   **Action:** Add a `custom_tgc_expand` entry in the YAML pointing to a template that skips API validation.

### 5. Argument Required, But No Definition Found (Encoder Issue)
*   **Symptom:** `The argument 'name' is required, but no definition was found`.
*   **Cause:** The field is dropped when TGC converts HCL to CAI JSON. Often the default Terraform provider's encoder drops it.
*   **Solution:** Add `tgc_ignore_terraform_encoder: true` to the field's definition in the YAML. This forces TGC to use its own expansion logic.

### 6. Schema Mismatch: CAI vs. GET API Response
*   **Symptom:** `At least 1 'feed_output_config' blocks are required` or `The argument 'config' is required`.
*   **Cause:** The JSON structure from CAI differs from the structure expected by the Terraform provider's schema or encoders.
*   **Solution:** Use `tgc_ignore_terraform_encoder: true` to bypass shared encoders and implement a custom TGC expander.

### 7. Ignored Terraform Decoder Needed
*   **Symptom:** Build error like `undefined: tpgresource.ParseImportId`.
*   **Cause:** Standard Terraform decoders contain logic (like API calls) that fails in the TGC environment (where no authorized client is available).
*   **Solution:** Add `tgc_ignore_terraform_decoder: true` to the resource or field in the YAML.

### 8. Missing Fields in HCL from CAI (Custom Flatten Issues)
*   **Symptom:** `missing fields in resource ... after cai2hcl conversion: [field_name]`.
*   **Cause:** Custom flatteners in Magic Modules often use `d.Get()`, which reads from state. During CAI -> HCL, there is no state, so `d.Get()` returns zero values.
*   **Solution:** Add `tgc_ignore_terraform_custom_flatten: true` to the field's definition in the YAML.

### 9. Field Only Exists in HCL, Not in CAI
*   **Symptom:** Test fails because it expects a field that CAI never provides.
*   **Solution:** Add `is_missing_in_cai: true` to the field in the resource YAML. This informs TGC that the field is not expected in the CAI input.

### 10. Incorrect CAI Asset Name Format
*   **Symptom:** A required argument (e.g., `'folder'`) is missing in the generated HCL.
*   **Cause:** TGC parses the CAI asset's `name` field to extract identifiers. If the format is wrong, it can't find the IDs.
*   **Solution:** Specify the correct format using `cai_asset_name_format` in the resource's YAML.
*   **Example:** `cai_asset_name_format: '//cloudasset.googleapis.com/folders/{{folder}}/feeds/{{feed_id}}'`.

### 11. Add Handwritten Subtests
*   **Symptom:** Message like `TestAccMonitoringAlertPolicy: test steps are unavailable`.
*   **Cause:** The main test doesn't exist, but subtests (e.g., `TestAccMonitoringAlertPolicy/basic`) do.
*   **Solution:** Add `tgc_tests` with the subtests list to the resource YAML file.

### 12. Undefined Function
*   **Symptom:** Message like `undefined: json`.
*   **Cause:** The `json` package is not imported in the generated code.
*   **Solution:** Add `"encoding/json"` to the import path in the `resource.go.tmpl` and ensure the import is preserved (e.g., using `_ = json.Unmarshal`).
