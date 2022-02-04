import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BestofTgComponent } from './bestof-tg.component';

describe('BestofTgComponent', () => {
  let component: BestofTgComponent;
  let fixture: ComponentFixture<BestofTgComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BestofTgComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BestofTgComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
