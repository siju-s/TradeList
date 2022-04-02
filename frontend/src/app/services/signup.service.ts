import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {environment} from "../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class SignupService {

  constructor(private http: HttpClient) {
  }

  signup(user: User) : Observable<any> {
    return this.http.post<any>(environment.gateway + '/signup', JSON.stringify(user))
  }

}

export interface User {
  Contact:Contact;
}

export interface Contact {
  FirstName:string;
  LastName:string;
  Email :string;
  Password :string;
}


