import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {environment} from "../../environments/environment";
import {Observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class LoginService {
  userFound : Boolean
  constructor(private http: HttpClient) {
  }

  login(login: Login) : Observable<any> {
    return this.http.post<any>(environment.gateway + '/login', JSON.stringify(login))
  }

}

export interface Login {
  Email :string;
  Password :string;
}
