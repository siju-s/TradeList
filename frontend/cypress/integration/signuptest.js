describe("Find the index page", function(){
    it("Finds the index page", function(){
        cy.visit('http://localhost:4200/')
    })
})


describe("Find the signup form", function(){
    it("Signup form working", function(){
        cy.contains('Register').click()
    
    })
})


describe("Fill the signup form", function(){
    it('Interacting with text fields',()=>{
       cy.get('input[name="firstname"]')
         .type('Mark')
    .should('have.value','Mark')
    .get('input[name="lastname"]')
    .type('Jacobs')
    .should('have.value','Jacobs')
    .get('input[id="register-email"]')
      .type('markjacobs@gmail.com')
 .should('have.value','markjacobs@gmail.com')
 .get('input[id="register-password"]')
 .type('12345678')
 .should('have.value','12345678')
    .get('button[class="btn btn-primary btn-block singup"]').click()
    })
   })

  