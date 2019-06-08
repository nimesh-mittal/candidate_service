# Candidate Service
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fnimesh-mittal%2Fcandidate_service.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fnimesh-mittal%2Fcandidate_service?ref=badge_shield)

sample go micro service

## How to build, install and run

##### `dep ensure`
#####  `go run main.go`

## Project Structure
### Three layers
##### Presentation Layer - Handlers
- This layer accepts request, parse request into Model objects and validates it
- It invokes Services to perform action related to request
- It accepts response from services, prepare JSON response and return to caller

##### Business Layer - Services
- This layer accepts parameters from resouce layer and perform actions
- To perform actions it calls DAO layer to fetch data from other APIs or Database
- It also perform business logic

##### Data Layer - DAO(Data access objects)
- This layer calls DB client functions or HTTP client functions to save/send or fetch/get data

#### Model
- A micro service usually manages lifecycle of one or more entities
- These entities can be represented as struct in GO under model package

##### Data Transfer Objects
- Most of the API calls does not need every field of model, some may need to pass metadata as well like limit, offset or errors
- Objects that represent input request content or output response content can be represent as structs and manager in model package with "dto" suffix

#### Infra
- Any infra related logic like setting up server, new-relic, swagger can be part of infra package

## Libraries
| Library | Purpose |
|---------|---------|
| Gin     | Router   |
| Logrus  | Logging  |
| GORM    | ORM for postgre, MySQL or other relational databases |
| Dep     | Dependency management |

## Todos
1. Docker **[done]**
2. Request validation **[done]**
3. Error handling [custom errors]
4. Redis support [In Progress]
5. Mongo support [In Progress]
6. Swagger support **[nested models]**
7. Log support **[flow id]**
8. New relic support
9. Heartbeat monitoring **[done]**
10. Code comments  **[minor fixes]**
11. Unit tests [In Progress]
12. Viper support **[done]**
13. Kinesis support **[Not Started]**
14. Rest call support **[Not Started]**


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fnimesh-mittal%2Fcandidate_service.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fnimesh-mittal%2Fcandidate_service?ref=badge_large)