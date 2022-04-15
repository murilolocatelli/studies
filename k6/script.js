import http from 'k6/http';
import { check, sleep } from 'k6';

export default function() {
  let res = http.get('http://f1a92123.ngrok.io/');
  check(res, {
    'status was 200': r => r.status == 200,
    'transaction time OK': r => r.timings.duration < 1000,
  });
  sleep(1);
}