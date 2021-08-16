import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Libro } from 'src/app/Modelo/Libro';
import { MyserviceService } from 'src/app/Service/myservice.service';

@Component({
  selector: 'app-add-libro',
  templateUrl: './add-libro.component.html',
  styleUrls: ['./add-libro.component.css']
})
export class AddLibroComponent implements OnInit {

  libro:Libro=new Libro();
  constructor(private router:Router, private service:MyserviceService) { }

  
  ngOnInit(): void {
  }
  
  Guardar(){
    this.service.createBook(this.libro)
    .subscribe(data => {
      alert("Se agrego con exito!")
      this.router.navigate(["listarLibro"]);
    })
  }

}
