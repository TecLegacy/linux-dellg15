# JWT basics

1. structure - header , payload , signature
2. header - contains 2 parts
   i. algorithm used to sign
   ii. type - type of token
3. payload - claims , 3 types of claims
   i. Registered claims - Registered claims include standard fields like issuer (iss), subject (sub), audience (aud), expiration time (exp), and issued at (iat).
   ii. Public claims
   iii. Private claims

## Goal of Application

1. RBAC - Senior , junior
   seniors - allowed to mutate add todo
   juniors - not allowed to add todo
