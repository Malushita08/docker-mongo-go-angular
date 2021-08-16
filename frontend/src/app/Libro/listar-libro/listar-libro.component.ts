import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Estudiante } from 'src/app/Modelo/Estudiante';
import { Libro } from 'src/app/Modelo/Libro';
import { MyserviceService } from 'src/app/Service/myservice.service';

@Component({
  selector: 'app-listar-libro',
  templateUrl: './listar-libro.component.html',
  styleUrls: ['./listar-libro.component.css']
})
export class ListarLibroComponent implements OnInit {

  libros: Libro[] = [];

  constructor(private myservice:MyserviceService, private router:Router){}

  ngOnInit(): void {
    this.myservice.getBooks()
    .subscribe(data => {
      this.libros = data.data.getBooks;
      console.log('data.data',this.libros);
    })
  }
  
  Remover(id:String){
    this.myservice.removeBook(id)
    .subscribe(data => {
      alert("Se elimino con exito!")
      this.router.navigate(["listarLibro"]);
    })
  }

  Editar(libro:Libro){
    localStorage.setItem("id",libro.id.toString());
    console.log('test',localStorage.getItem("id"));
    this.router.navigate(['editLibro']);
  }



}
