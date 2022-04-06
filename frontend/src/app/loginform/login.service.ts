import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {environment} from "../../environments/environment";
import {Observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class LoginService {
  userFound : Boolean
  user : User
  constructor(private http: HttpClient) {
  }

  login(login: Login) : Observable<any> {
    return this.http.post<any>(environment.gateway + '/login', JSON.stringify(login))
  }

  signup(user: User) : Observable<any> {
    return this.http.post<any>(environment.gateway + '/signup', JSON.stringify(user))
  }

  setUser(user: User) {
    this.user = user
  }

  getUser() : User {
    return this.user
  }

  getUserId() : number {
    if (this.user == null) {
        return 0
    }
    const id = this.user.ID
    if (id == null) {
      return 0
    }
    return id
  }

}

export interface Login {
  Email :string;
  Password :string;
}

export interface User {
  ID?:number;
  Contact:Contact;
}

export interface Contact {
  FirstName:string;
  LastName:string;
  Email :string;
  Password :string;
}
