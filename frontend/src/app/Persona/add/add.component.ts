import { Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';
import { Estudiante } from 'src/app/Modelo/Estudiante';
import { MyserviceService } from 'src/app/Service/myservice.service';

@Component({
  selector: 'app-add',
  templateUrl: './add.component.html',
  styleUrls: ['./add.component.css']
})
export class AddComponent implements OnInit {

  estudiante:Estudiante=new Estudiante();
  constructor(private router:Router, private service:MyserviceService) { }

  
  ngOnInit(): void {
    
  }

  /*
  Guardar(){
    alert("Se agrego con exito!")
    this.router.navigate(["listar"]);
  }
  */
  
  Guardar(){
    this.service.createStudent(this.estudiante)
    .subscribe(data => {
      alert("Se agrego con exito!")
      this.router.navigate(["listar"]);
    })
  }

}
