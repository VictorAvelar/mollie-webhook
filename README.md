# Mollie webhook

A mollie webbook handler that fits 80% of the use cases and avoids the need to code stuff on your own.

## Why?

Mollie native webhook implementation hits your provided url with an ID of the resource you want updates for, this requires
your to call their API to retrieve the actual information needed to perform any operation.

Mollie does this based on security concerns with sending sensitive date (such as the one contained in a payment) to a random
URL on the internet, requiring you to call their API ensures that you have a valid token and that you have the correct permissions
to access the resource.

## Solution

This package does exactly what you would expect, it wraps the code and handlers required to do what you need from a webhook.

If it receives a payment id it will return the payment object, and so on for each of Mollie's resources.
