# Vulnerability Management Basics

This document explains several important cybersecurity concepts in
simple terms:

-   NVD
-   SCAP
-   CVSS
-   SBOM
-   VEX

These standards help security tools **identify, understand, and
prioritize software vulnerabilities**.

------------------------------------------------------------------------

# 1. National Vulnerability Database (NVD)

The **National Vulnerability Database (NVD)** is a public database that
stores information about known security vulnerabilities.

It is maintained by the U.S. government and used by many cybersecurity
tools and vulnerability scanners.

## What the NVD Contains

The NVD includes:

-   Information about known vulnerabilities
-   Links to security advisories and patches
-   Standard identifiers for affected software
-   Severity scores for vulnerabilities

## Why the NVD Is Important

Security tools use the NVD to:

-   Detect vulnerable software
-   Understand how dangerous a vulnerability is
-   Decide which issues should be patched first

------------------------------------------------------------------------

# 2. Security Content Automation Protocol (SCAP)

**SCAP** is a collection of standards used to automate security checks
and vulnerability management.

In simple terms:

SCAP allows different security tools to **share security information in
a standardized format**.

## What SCAP Is Used For

SCAP helps automate:

-   Vulnerability scanning
-   Security configuration checks
-   Compliance monitoring
-   Security reporting

## Important SCAP Components

### CVE (Common Vulnerabilities and Exposures)

A global list of publicly known vulnerabilities.

Each vulnerability receives a unique ID.

Example: CVE-2021-44228

### CPE (Common Platform Enumeration)

A standardized naming system for software and hardware.

This helps security tools identify exactly which product and version is
affected.

### OVAL (Open Vulnerability and Assessment Language)

A language used to describe how systems should be tested for
vulnerabilities.

### XCCDF (Extensible Configuration Checklist Description Format)

A format used for defining security configuration checklists and
compliance rules.

------------------------------------------------------------------------

# 3. Common Vulnerability Scoring System (CVSS)

**CVSS** is a system used to measure the severity of a vulnerability.

Each vulnerability receives a score between **0.0 and 10.0**.

The higher the score, the more severe the vulnerability.

## CVSS Severity Levels

  Score         Severity
  ------------- ----------
  0.0 -- 3.9    Low
  4.0 -- 6.9    Medium
  7.0 -- 8.9    High
  9.0 -- 10.0   Critical

## How CVSS Scores Are Calculated

CVSS looks at factors such as:

-   How easy the vulnerability is to exploit
-   Whether special permissions are required
-   Whether user interaction is needed
-   The impact on confidentiality, integrity, and availability

These factors are combined to calculate the final score.

------------------------------------------------------------------------

# 4. Software Bill of Materials (SBOM)

A **Software Bill of Materials (SBOM)** is a list of all components that
make up a piece of software.

You can think of it like an **ingredient list for software**.

It shows:

-   Libraries used in the software
-   Open‑source dependencies
-   Versions of components
-   The relationships between components

## Why SBOMs Are Important

SBOMs help organizations:

-   Understand what software components they are using
-   Identify vulnerable dependencies
-   Respond quickly when new vulnerabilities are discovered
-   Improve software supply chain security

For example:

If a vulnerability is discovered in a library like Log4j, an SBOM can
help quickly determine whether your software includes that component.

------------------------------------------------------------------------

# 5. Vulnerability Exploitability eXchange (VEX)

**VEX** is a document that explains whether a vulnerability actually
affects a specific product.

Sometimes a vulnerability exists in a component, but the product using
that component is **not actually vulnerable**.

VEX helps communicate this clearly.

## What VEX Answers

A VEX document can indicate:

-   The vulnerability affects the product
-   The vulnerability does NOT affect the product
-   The vulnerability is under investigation
-   The vulnerability has already been fixed

## Why VEX Is Useful

Without VEX:

Security teams may panic when a vulnerability appears in a dependency.

With VEX:

Vendors can explain whether their software is truly affected.

This reduces:

-   False alarms
-   Unnecessary patching
-   Investigation time

# 5. Vulnerability Exploitability eXchange (VEX)

**VEX** documents explain whether a vulnerability actually affects a product.

Sometimes a vulnerable component exists, but the product **is not affected**.

VEX communicates this clearly.

## What VEX Answers

A VEX document indicates if a vulnerability:

- Affects the product
- Does NOT affect the product
- Has been fixed
- Is under investigation


### VEX Document Formats

  1. CSAF (Common Security Advisory Framework) is a JSON-based standard for publishing security advisories.
  2. CycloneDX is a SBOM standard. Embeds VEX statements directly in the SBOM. Links vulnerabilities to components in the software.
  3. SPDX (Software Package Data Exchange) is another SBOM standard widely used in open-source communities. VEX statements attach vulnerability status to SPDX component identifiers.


## Common VEX Status Values

| Status | Meaning |
|------|------|
| affected | The product is vulnerable |
| not_affected | The product is not vulnerable |
| fixed | The vulnerability existed but is patched |
| under_investigation | Still analyzing impact |

## Key Fields in a VEX Statement

- **Product**: The software being evaluated  
  Example: `ExampleApp 2.0`
- **Vulnerability**: Usually a CVE ID  
  Example: `CVE-2021-44228`
- **Status**: Affects or not  
  Example: `not_affected`
- **Justification**: Why the product is safe  
  Examples: `component_not_present`, `vulnerable_code_not_in_execute_path`
- **Impact Statement (Optional)**: Extra explanation  
  Example: `Log4j library present but vulnerable class not used`


### Example VEX (JSON)

```json
{
  "product": "ExampleApp 2.0",
  "vulnerability": "CVE-2021-44228",
  "status": "not_affected",
  "justification": "vulnerable_code_not_in_execute_path",
  "impact_statement": "Log4j library present but vulnerable functionality not used."
}

------------------------------------------------------------------------

# Important Note About Risk

A vulnerability score (like CVSS) shows **severity**, but it does not
always represent the real risk to a specific organization.

Security teams should also consider:

-   Asset importance
-   Whether an exploit exists
-   Network exposure
-   Business impact

------------------------------------------------------------------------

# Quick Summary

  -----------------------------------------------------------------------
  Concept                             Description
  ----------------------------------- -----------------------------------
  NVD                                 Public database of known
                                      vulnerabilities

  SCAP                                Standards for automating
                                      vulnerability and compliance checks

  CVSS                                Scoring system used to measure
                                      vulnerability severity

  SBOM                                A list of all software components
                                      and dependencies

  VEX                                 A document explaining whether a
                                      vulnerability affects a product
  -----------------------------------------------------------------------

Together, these concepts help organizations **understand vulnerabilities
and manage security risks more effectively**.
