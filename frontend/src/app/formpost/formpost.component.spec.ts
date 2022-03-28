import { ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';

import { FormpostComponent } from './formpost.component';

describe('FormpostComponent', () => {

  let component: FormpostComponent;
  let fixture: ComponentFixture<FormpostComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FormpostComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FormpostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });
  
  // it('should create', () => {
  //   expect(component.numberofLocations).toBe(4);
  // });
  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should create ', () => {
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('#formtitle').textContent).toBe('Create a Post!');
    
  });

});
