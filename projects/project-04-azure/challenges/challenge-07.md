# Challenge 7: Secure an application using identity, authentication, authorization, least privilege, and managed identities where supported

## Difficulty
Intermediate–Advanced

## Problem statement and user story
The booking platform contains sensitive operational data, including bookings, customer contact metadata, and internal business rules. The app must authenticate people correctly and restrict service-to-service access to only the permissions needed.

## Practical outcome
Participants secure an app using identity, access control, secrets handling, and least privilege patterns.

## Learning objectives
- distinguish authentication from authorization
- use identity-based access patterns in Azure
- reason about least privilege and managed identity
- explain why secrets should not be hardcoded or stored in source control

## Prerequisites and setup
- deployed app or service in Azure
- understanding of app roles or user roles
- access to a test identity or environment

## Azure services under consideration
- Microsoft Entra ID
- Azure RBAC
- Azure Key Vault
- App Service or Azure Functions managed identity support

### Verified free-tier / cost notes
- many identity and access features are included as part of Azure tenant capabilities, but feature tiers and pricing vary.
- Azure Key Vault usage is billable based on transactions and other consumption factors.
- Use current pricing and eligibility guidance before deployment.

## Reference architecture and request/data flow
```text
User -> Identity provider -> app -> authorization checks -> protected data
Service -> managed identity -> Key Vault or storage -> restricted access
```

## Functional and non-functional requirements
Functional:
- only authorized users can access protected operations
- service-to-service access is scoped and limited
- secrets are not visible in repo files or logs

Non-functional:
- least privilege by default
- clear separation between admin and user roles
- reliable and auditable access model

## Acceptance criteria
- valid users can access allowed resources
- invalid or unauthorized users are rejected
- app does not contain hardcoded secrets
- service identity is used where supported

## Minimal implementation path
1. identify protected operations and roles
2. configure authentication for the app
3. grant narrow permissions to roles or identities
4. move secrets to Key Vault
5. test access as authorized and unauthorized users

## More advanced implementation path
- add app roles or scopes
- implement conditional or group-based access rules
- use managed identity for a worker or storage access path

## Stretch goal
- integrate with a real enterprise identity model
- add auditing or access logs
- define a security review checklist for future apps

## Test plan
- valid user can access protected route
- unauthorized user is denied
- service identity access is restricted to required resources
- secret retrieval from Key Vault succeeds only with permissions

### Negative tests and failure cases
- invalid token
- user missing required role
- secret leaked in logs
- over-privileged role assignment

## Security and privacy considerations
- protect sensitive customer data
- minimize data stored in logs
- prefer managed identities over long-lived secrets
- review role assignments regularly

## Cost risks, usage limits, and teardown checklist
- Key Vault operations and monitoring can add cost
- extra logging or audit volume increases cost
- review role assignments and remove test identities when done

## AI-agent prompts
- “Review this booking app architecture for identity and least privilege gaps. What would be the minimum secure baseline?”
- “Explain the difference between app authentication, user authentication, and service-to-service authorization in Azure.”

## Questions participants must answer
- Why is a secret in source control a problem?
- What is a managed identity and why is it valuable?
- What is the difference between authentication and authorization?

## Deliverables for the repository
- auth model documentation
- access matrix or role summary
- configuration examples without real secrets
- security checklist

## Reflection exercise and retrospective questions
- What if a developer is accidentally granted too much access?
- Which permissions are essential for the app and which are unnecessary?

## Rubric
- identity model correctness: 30%
- least privilege reasoning: 25%
- secret handling: 20%
- testing and validation: 15%
- documentation: 10%

## Hint 1
Start by identifying the trust boundaries and the actors in the system.

## Hint 2
Prefer narrow role assignments over broad admin permissions.

## Hint 3
If a service can use managed identity, prefer it over a stored secret.
