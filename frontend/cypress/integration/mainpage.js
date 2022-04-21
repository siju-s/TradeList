describe("Find the index page", function(){
    it("Finds the index page", function(){
        cy.visit('http://localhost:4200')
    })
})

describe("Interacts with search bar", function(){
    it('Interacting with search bar',()=>{
        cy.get('input[name="searchbar"]').click()
        .type("deve")
   .should('have.value','deve')
      
    })
   })

   describe("Interacts with list and grid view bar", function(){
    it('Interacting with list and grid view bar',()=>{
        cy.contains('List').click()
      
    })
   })

   describe("Interacts with list and grid view bar", function(){
    it('Interacting with list and grid view bar',()=>{
        cy.contains('Grid').click()
        cy.pause()
    })
   })




describe("Interacts with navbar", function(){
    it('Interacting with navbar',()=>{
       cy.contains('Jobs').click()
      
    })
   })
   
   describe("Interacts with navbar", function(){
    it('Interacting with navbar',()=>{
       cy.contains('Property').click()
      
    })
   })

   describe("Interacts with navbar", function(){
    it('Interacting with navbar',()=>{
       cy.contains('For Sale').click()
      
    })
   })

   describe("Interacts with navbar", function(){
    it('Interacting with navbar',()=>{
       cy.contains('Services').click()
      
    })
   })

   describe("Interacts with navbar", function(){
    it('Interacting with navbar',()=>{
       cy.contains('Community').click()
      
    })
   })
   
  