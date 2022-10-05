[]: # Title: A simple project using GraphQL and gqlgen
[]: # Date: 2022-10-05
[]: # Tags: graphql, gqlgen, go
[]: # Category: go
[]: # Slug: a-simple-project-using-graphql-and-gqlgen
[]: # Description: A simple project using GraphQL and gqlgen

[]: # An example:

[]: ## mutation createCategiry {
[]: ##      createCategory(input: {
[]: ##          name: "Developer",
[]: ##          description: "Courses to developers"
[]: ##      }
[]: ##   ){
[]: ##          id
[]: ##          name
[]: ##          description
[]: ##     }
[]: ## }
[]: ## 
[]: ## query getCategories {
[]: ##      categories {
[]: ##          id
[]: ##          name
[]: ##          description
[]: ##      }
[]: ## }
