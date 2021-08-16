
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MyserviceService } from './Service/myservice.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'Servicio Biblioteca';
  books = [];

  constructor(private router:Router, private myservice:MyserviceService){}
  //constructor(private myservice:MyserviceService){}

  ngOnInit(): void{
    //this.getBooks();
  }
  
  ListarEstudiantes(){
    this.router.navigate(['listar']);
  }
  NuevoEstudiante(){
    this.router.navigate(['add']);
  }
  ListarLibros(){
    this.router.navigate(['listarLibro']);
  }
  NuevoLibro(){
    this.router.navigate(['addLibro']);
  }
  
  /*getBooks(){
    this.myservice.getBooks().subscribe((data) => {
      this.books = data.data.getBooks;
      console.log(this.books);
    })
  }*/

}
