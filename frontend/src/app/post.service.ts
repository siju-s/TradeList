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

  createPost(post: Post, files?: FileList) {
    const formData = new FormData();
    formData.append('data', JSON.stringify(post))
    for (let i = 0; i < files!.length; i++) {
      formData.append('files', files![i], files![i].name);
    }
    const headers = new HttpHeaders()
      .append("Content-Type", "application/json")
      .append("Access-Control-Allow-Origin", "*")
      .append("Accept", "multipart/form-data");

    console.log(formData)
    this.http.post<any>(environment.gateway + '/post', formData).subscribe(data => {
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
  Image?:FileList;
}

export interface Response {
   data: Array<Post>
   message: string
   status : number
}
