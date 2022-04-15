import http from 'k6/http';
import { check, sleep } from 'k6';

export default function() {
  let res = http.post('http://7e9eabc9.ngrok.io/v1/insert?nome=nome&email=email');
  check(res, {
    'status was 201': r => r.status == 201,
    'transaction time OK': r => r.timings.duration < 1000,
  });
  sleep(1);
} 