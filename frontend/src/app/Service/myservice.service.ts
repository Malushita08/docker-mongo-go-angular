import { Injectable } from '@angular/core';
//import { Observable } from '@apollo/client/utilities';
import { Observable } from 'rxjs';
import { Apollo, gql } from 'apollo-angular';
import { Estudiante } from '../Modelo/Estudiante';
import { Libro } from '../Modelo/Libro';
import { ListarLibroComponent } from '../Libro/listar-libro/listar-libro.component';


const getBooksQuery = gql`
  query {
    getBooks{
      author
      id
      price
      tittle
    }
  }
`;
/*
const getBookQuery = gql`
  query (getBook($id:String!)){
    getBook(id:$id){
      id
      tittle
      author
      price
    }
  }
`;*/

const getStudentsQuery = gql`
  query {
    getStudents{
      code
      dni
      id
      lastName
      name
    }
  }
`;

const addStudent = gql`
  mutation addStudent($name:String!, $lastName:String!, $dni:String!, $code:String!){
    addStudent(
      student:{
        name:$name
        lastName:$lastName
        dni: $dni
        code: $code
      }){
      id
      name
      dni
    }
  }
`;

const addBook = gql`
  mutation addBooks($tittle:String!, $price:Float!, $author:String!){
    addBooks(
      book:{
        tittle:$tittle
        price:$price
        author:$author
    }){
      id
      tittle
      price
    }
  }
`;

const removeStudent = gql`
  mutation removeStudent($id: String!){
    removeStudent(
      id: $id
    ){
      id
    }
  }
`;

const removeBook = gql`
  mutation removeBook($id: String!){
    removeBook(
      id: $id
    ){
      id
    }
  }
`;

const updateBook = gql`
  mutation updateBooks($id:String!, $author:String!, $price:Float!, $tittle:String!){
    updateBooks(
      id: $id
      book:{
        author: $author
        price: $price
        tittle: $tittle
    }){
      id
    }
  }
`;



@Injectable({
  providedIn: 'root'
})
export class MyserviceService {

  constructor(private apollo:Apollo) { }

  getBooks(): Observable<any>{
    return this.apollo.watchQuery<any>({
      query: getBooksQuery,
    }).valueChanges;    
  }
/*
  getBook(id:string): Observable<any>{
    return this.apollo.watchQuery<any>({
      query: getBookQuery,
      variables:{
        id:id
      }
    }).valueChanges;    
  }*/

  getStudents(): Observable<any>{
    return this.apollo.watchQuery<any>({
      query: getStudentsQuery,
    }).valueChanges;    
  }
  createStudent(estudiante:Estudiante){
    return this.apollo.mutate({
      mutation: addStudent,
      variables:{
        name: estudiante.name,
        lastName: estudiante.lastName,
        dni: estudiante.dni,
        code: estudiante.code
      }
    });
  }
  createBook(libro:Libro){
    return this.apollo.mutate({
      mutation: addBook,
      variables:{
        tittle:libro.tittle,
        price:libro.price,
        author:libro.author
      }
    });
  }
  removeStudent(id:String){
    return this.apollo.mutate({
      mutation: removeStudent,
      variables:{
        id:id
      }
    });
  }
  removeBook(id:String){
    return this.apollo.mutate({
      mutation: removeBook,
      variables:{
        id:id
      }
    });
  }
  updateBook(libro:Libro){
    return this.apollo.mutate({
      mutation: updateBook,
      variables:{
        id:libro.id,
        author:libro.author,        
        price:libro.price,
        tittle:libro.tittle        
      }
    });
  }
}
