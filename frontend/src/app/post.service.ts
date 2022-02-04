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

  getPosts(): Observable<Array<Post>> {
    const result = this.http.get<Array<Post>>(environment.gateway + '/post');
    console.log("GetPosts", result)
    // result.subscribe(data => {
    //   console.log(data)
    // })
    return result
  }

  createPost(post: Post) {
    const headers = new HttpHeaders().append("Content-Type", "application/json")
      .append("Access-Control-Allow-Origin", "*");

    console.log("Post is:", post.description)
    this.http.post<Post>(environment.gateway + '/post', JSON.stringify(post), {headers: headers}).subscribe(data => {
      console.log(data)
    });
  }
}


export interface Post {
  sellerid: number;
  categoryid: number;
  subcategoryid: number;
  Title: string;
  description: string;
  CreatedAt?:string;
}
