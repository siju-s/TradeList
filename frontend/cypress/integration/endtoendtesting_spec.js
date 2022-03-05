describe("Test 1", function(){
    it("Finds the index page", function(){
        cy.visit('http://localhost:4200/')
    })
})

describe("Test 2", function(){
    it("Login Form working", function(){
        cy.contains('Login').click()
    })
})

describe("Test 3", function(){
    it("Signup form working", function(){
        cy.contains('Signup').click()
    })
})


describe("Test 4", function(){
    it("Find checkbox in create post form is  functional", function(){
        cy.contains('Show my personal information').click({force: true})
       
    })

})


describe("Test 5", function(){
    it("Checks if navbar is  functional", function(){
        cy.contains('Community').click()
        cy.contains('Housing').click( {force: true})
        cy.contains('Gigs').click({force: true})
        cy.contains('Jobs').click({force: true})
       
    })
})


describe("Test 6", function(){
    it("Checks if navbar Services is  functional", function(){
        cy.contains('Services').click({force: true})
       
    })

})
describe("Test 7", function(){
    it("Checks if navbar Forums is  functional", function(){
        cy.contains('Forums').click({force: true})
       
    })

})
describe("Test 8", function(){
    it("Checks if navbar For Sale is  functional", function(){
        cy.contains('For Sale').click({force: true})
       
    })

})

describe("Test 9", function(){
    it("Checks if Grid component navigation is functional", function(){
        cy.contains('Household').click({force: true})
       
    })

})

