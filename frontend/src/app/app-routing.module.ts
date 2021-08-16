import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ListarComponent } from './Persona/listar/listar.component';
import { AddComponent } from './Persona/add/add.component';
import { EditComponent } from './Persona/edit/edit.component';
import { ListarLibroComponent } from './Libro/listar-libro/listar-libro.component';
import { AddLibroComponent } from './Libro/add-libro/add-libro.component';
import { EditLibroComponent } from './Libro/edit-libro/edit-libro.component';
import { AppComponent } from './app.component';

const routes: Routes = [
  
  {path:'listar', component:ListarComponent},
  {path:'add',component:AddComponent},
  {path:'edit',component:EditComponent},
  {path:'listarLibro', component:ListarLibroComponent},
  {path:'addLibro', component:AddLibroComponent},
  {path:'editLibro', component:EditLibroComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
