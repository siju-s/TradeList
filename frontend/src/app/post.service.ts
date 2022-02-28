import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";
import {environment} from "../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class PostService {

  constructor(private http: HttpClient) {
  }

  getPosts(): Observable<Response> {
    const result = this.http.get<Response>(environment.gateway + '/post');
    console.log("GetPosts", result)
    // result.subscribe(data => {
    //   console.log(data)
    // })
    return result
  }

  createPost(post: Post) {
    const headers = new HttpHeaders().append("Content-Type", "application/json")
      .append("Access-Control-Allow-Origin", "*");

    console.log("Post is:", post.Description)
    this.http.post<Post>(environment.gateway + '/post', JSON.stringify(post), {headers: headers}).subscribe(data => {
      console.log(data)
    });
  }
}


export interface Post {
  Sellerid: number;
  Categoryid: number;
  Subcategoryid: number;
  Title: string;
  Description: string;
  CreatedAt?:string;
}

export interface Response {
   data: Array<Post>
   message: string
   status : number
}
