## Rate-Limited Notification Service

We have a Notification system that sends out email notifications of various types (update, daily news, project invitations, etc.).  
We need to protect recipients from getting too many emails, either due to system errors or due to abuse, so let's limit the number of emails sent to them by implementing a rate-limited version of NotificationService.  
The system must reject requests that are over the limit.

Some sample notification types and rate limit rules, e.g.:
- Status: not more than 2 per minute for each recipient
- News: not more than 1 per day for each recipient
- Marketing: not more than 3 per hour for each recipient

### Solution

This solution is based on the Leaky Bucket Algorithm provided by `golang.org/x/time/rate` package.  
As notifications are sent to the recipients, the tokens (rate) are decremented by one until there are none available and the Allow() method returns false. Tokens are refilled at rate.  
Since we have to control the rate of different types of notifications for each recipient, we have to create `rate.Limiter` with that combination and store them in a map.