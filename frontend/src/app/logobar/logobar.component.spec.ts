import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LogobarComponent } from './logobar.component';

describe('LogobarComponent', () => {
  let component: LogobarComponent;
  let fixture: ComponentFixture<LogobarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LogobarComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LogobarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
