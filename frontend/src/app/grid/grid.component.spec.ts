// import { ComponentFixture, TestBed } from '@angular/core/testing';

// import { GridComponent } from './grid.component';

// describe('GridComponent', () => {
//   let component: GridComponent;
//   let fixture: ComponentFixture<GridComponent>;

//   beforeEach(async () => {
//     await TestBed.configureTestingModule({
//       declarations: [ GridComponent ]
//     })
//     .compileComponents();
//   });

//   beforeEach(() => {
//     fixture = TestBed.createComponent(GridComponent);
//     component = fixture.componentInstance;
//     fixture.detectChanges();
//   });

//   it('should create', () => {
//     expect(component).toBeTruthy();
//   });
// });

// #############

import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GridComponent } from './grid.component';

describe('GridComponent', () => {
  let component: GridComponent;
  let fixture: ComponentFixture<GridComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GridComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(GridComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should create', () => {
    expect(component.componentName).toBe("user");
  });
  it('should create', () => {
    expect(component.title).toBe("Card View Demo");
  });
  it('should create', () => {
    expect(component.gridColumns).toBe(3);
  });
  it('should create', () => {
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('#check').textContent).toBe('Housing options');
    
  });


});
