import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListarLibroComponent } from './listar-libro.component';

describe('ListarLibroComponent', () => {
  let component: ListarLibroComponent;
  let fixture: ComponentFixture<ListarLibroComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListarLibroComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ListarLibroComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
