# CacheGO

It's an exercise to solve a challenge that needs to create a library that handles the cache requests to optimize the database queries when the user data info was exposed previously.

For example: If we have 1000 requests for user data info and 100 of them were sent to the client before, we need to take those 100 from the cache strategy instead of requesting again from the database.

# An alternative solution

A well structured solution [here](https://github.com/chrobson/RedisCache)