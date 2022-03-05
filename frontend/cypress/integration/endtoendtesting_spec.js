describe("Test 1", function(){
    it("Finds the index page", function(){
        cy.visit('http://localhost:4200/')
    })
})

describe("Test 2", function(){
    it("Finds the login button", function(){
        cy.contains('Login').click()
    })
})

describe("Test 3", function(){
    it("Finds the signup button", function(){
        cy.contains('Signup').click()
    })
})



describe("Test 11", function(){
    it("Find description in create post form is  functional", function(){
        cy.contains('Description').click({force: true})
       
    })

})

describe("Test 12", function(){
    it("Find address in create post form is  functional", function(){
        cy.contains('Address').click({force: true})
       
    })

})

describe("Test 13", function(){
    it("Find checkbox in create post form is  functional", function(){
        cy.contains('Show my personal information').click({force: true})
       
    })

})


describe("Test 4", function(){
    it("Checks if navbar is  functional", function(){
        cy.contains('Community').click()
       
    })
})
describe("Test 5", function(){
    it("Checks if navbar is functional", function(){
        cy.contains('Housing').click( {force: true})
       
    })

})
describe("Test 6", function(){
    it("Checks if navbar is  functional", function(){
        cy.contains('Gigs').click({force: true})
       
    })

})
describe("Test 7", function(){
    it("Checks if navbar is functional", function(){
        cy.contains('Jobs').click({force: true})
       
    })

})
describe("Test 8", function(){
    it("Checks if navbar is  functional", function(){
        cy.contains('Services').click({force: true})
       
    })

})
describe("Test 9", function(){
    it("Checks if navbar is  functional", function(){
        cy.contains('Forums').click({force: true})
       
    })

})
describe("Test 10", function(){
    it("Checks if navbar is  functional", function(){
        cy.contains('For Sale').click({force: true})
       
    })

})

