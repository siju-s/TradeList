describe("Find the index page", function(){
    it("Finds the index page", function(){
        cy.visit('http://localhost:4200')
    })
})

describe("Find the login form", function(){
    it("Login Form working", function(){
        cy.contains('Login').click()
    })
})

describe("Fill the login form", function(){
 it('Interacting with text fields',()=>{
    cy.get('input[name="email"]')
      .type('janethomas@gmail.com')
 .should('have.value','janethomas@gmail.com')
 .get('input[name="password"]')
 .type('adamsand')
 .should('have.value','adamsand')
 cy. pause() 
//  .get('button[class="btn btn-primary btn-block login"]').click()

 })
})

describe("Find the create a post form", function(){
    it("Create a post form working", function(){
        cy.contains('Create a Post').click()
    })
})

describe("Fill the Create a post form", function(){
    it('Interacting with text fields',()=>{
       cy.get('input[formControlName="Title"]')
         .type('Test Title')
    .should('have.value','Test Title')
    .get('input[formControlName="Price"]')
         .type('Test Price')
    .should('have.value','Test Price')
    .get('input[formControlName="email"]')
      .type('test@gmail.com')
 .should('have.value','test@gmail.com')
 .get('input[formControlName="PhoneNo"]')
 .type('12345678')
 .should('have.value','12345678')
 cy.contains('Save').click()
    })
   })