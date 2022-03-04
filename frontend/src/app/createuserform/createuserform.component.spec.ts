import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateuserformComponent } from './createuserform.component';

describe('CreateuserformComponent', () => {
  let component: CreateuserformComponent;
  let fixture: ComponentFixture<CreateuserformComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateuserformComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateuserformComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
