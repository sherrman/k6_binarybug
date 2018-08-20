import http from "k6/http";
import { check, group } from "k6";

const rand1 = open("./random1.bin", "b");

export default function() {
    let binaryFromPythonHttp = `http://127.0.0.1:8000/random1.bin`;
    let getBinaryUrl = `http://localhost:9999/binary?f=random1.bin`;
    let verifyBinaryUrl = `http://localhost:9999/compare?f=random1.bin`;

    let binFromHttpResponse = http.get(getBinaryUrl);
    let binFromPythonResponse = http.get(binaryFromPythonHttp);

    if (binFromHttpResponse.status != 200) {
        fail("unable to get binary from test server");
        return;
    }

    let binaryData = [binFromPythonResponse.body, rand1, binFromHttpResponse.body];

    binaryData.forEach(function (d) {
        console.log(`Data length: ${d.length}`);
        let compareResponse = http.post(verifyBinaryUrl, d);
        let statusCode = compareResponse.status;

        check(compareResponse, {
            "response code was 200": (res) => res.status == 200,
            "failed compare 400": (res) => res.status == 400,
        });
    });
}
