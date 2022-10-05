# A simple project using GraphQL and gqlgen

## An example:

 mutation createCategory {
      createCategory(input: {
          name: "Developer",
          description: "Courses to developers"
      }
   ){
          id
          name
          description
     }
 }
 
 query getCategories {
      categories {
          id
          name
          description
      }
 }
