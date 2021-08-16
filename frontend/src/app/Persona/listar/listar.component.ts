import { Component, OnInit } from '@angular/core';
import { MyserviceService } from 'src/app/Service/myservice.service';
import { Estudiante } from 'src/app/Modelo/Estudiante';
import { Router } from '@angular/router';

@Component({
  selector: 'app-listar',
  templateUrl: './listar.component.html',
  styleUrls: ['./listar.component.css']
})
export class ListarComponent implements OnInit {

  estudiantes: Estudiante[] = [];

  constructor(private myservice:MyserviceService, private router:Router){}

  ngOnInit(): void {
    this.myservice.getStudents()
    .subscribe(data => {
      this.estudiantes = data.data.getStudents;
      console.log('data.data',this.estudiantes);
    })
  }

  Remover(id:String){
    this.myservice.removeStudent(id)
    .subscribe(data => {
      alert("Se elimino con exito!")
      this.router.navigate(["listar"]);
    })
  }

}
