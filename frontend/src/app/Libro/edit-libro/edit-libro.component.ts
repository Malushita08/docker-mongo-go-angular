import { stringify } from '@angular/compiler/src/util';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Libro } from 'src/app/Modelo/Libro';
import { MyserviceService } from 'src/app/Service/myservice.service';

@Component({
  selector: 'app-edit-libro',
  templateUrl: './edit-libro.component.html',
  styleUrls: ['./edit-libro.component.css']
})
export class EditLibroComponent implements OnInit {

  libro:Libro=new Libro();
  constructor(private router:Router, private service:MyserviceService) { }

  
  ngOnInit(): void {
    
    this.Editar();
  }
  /*
  Guardar(){
    this.service.updateBook(this.libro)
    .subscribe(data => {
      alert("Se agrego con exito!")
      this.router.navigate(["listarLibro"]);
    })
  }*/


  Editar(){
    
    console.log('id',localStorage.getItem("id"));
    console.log('metaaaa');

/*
    this.service.getBook(localStorage.getItem("id") || "")
    .subscribe(data => {
      //this.libro = data.data.getBooks;
      console.log('data.data',data);
    })
    //console.log('editar');
    */
  }

  Actualizar(){
    
    alert("En desarrollo!")
/*
    this.service.updateBook(this.libro)
    .subscribe(data => {
      //this.libro=data;
      alert("Se actualizo con exito!")
      this.router.navigate(["listarLibro"]);
    })*/
  }

}
