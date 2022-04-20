describe("Find the index page", function(){
    it("Finds the index page", function(){
        cy.visit('http://localhost:4200/')
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
 .get('button[class="btn btn-primary btn-block login"]').click()
 })
})

