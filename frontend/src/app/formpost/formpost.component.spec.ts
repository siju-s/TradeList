import { async, ComponentFixture, TestBed } from '@angular/core/testing';
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
  it(`should have as title 'angular-unit-test'`, async(() => {
    const fixture = TestBed.createComponent(FormpostComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('angular-unit-test');
  }));
  it('should render title in a p tag', async(() => {
    const fixture = TestBed.createComponent(FormpostComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('p').textContent).toContain('>Upload images related to your posting');
  }));
  it('should create', () => {
    expect(component.selectedCategoryId).toBe(0);
  });

});
