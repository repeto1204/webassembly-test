import React, {useEffect, useState} from 'react';
import { instantiateStreaming } from "assemblyscript/lib/loader";

const App = () => {
    let WASM: any;

    // ========================================================================= //
    // RUST
    // const callRust = async () => {
    //     WASM = await import('../../public/rswasm.js');
    //     WASM.greet('WORLD!');
    // }
    // callRust();
    // ========================================================================= //

    // ========================================================================= //
    // ASSEMBLYSCRIPT
    // const callAs = async () => {
    //     WASM = await instantiateStreaming(fetch("./index.wasm"), {});
    //     alert(WASM.say());
    //     alert(WASM.__getString(WASM.say()));
    // };
    // const runAs = () => {
    //     console.log(WASM.__getString(WASM.say()));
    // }
    // callAs();
    // ========================================================================= //

    // ========================================================================= //
    // GO
    const [plainText, setPlainText] = useState('');
    const [cipherText, setCipherText] = useState('');
    const [decryptText, setDecryptText] = useState('');

    useEffect( () => {
        callGo();
    }, [])
    const callGo = async () => {
        const go = new Go();
        const WASM = await WebAssembly.instantiateStreaming(fetch("./index.wasm"), go.importObject);
        go.run(WASM.instance);
    };
    // ========================================================================= //



    return (
        <>
            <div>GO</div>
            {/*RUST*/}
            {/*<button onClick={() => {WASM.run()}}>CLICK!</button>*/}
            {/*AssemblyScript*/}
            {/*<button onClick={() => {runAs()}}>CLICK!</button>*/}
            {/*GO*/}
            <input type="text" onChange={(e) => setPlainText(e.target.value)}/>
            <button onClick={() => setCipherText(WebassemblyEncrypt(plainText))}>ENCRYPT</button>
            <button onClick={() => setDecryptText(WebassemblyDecrypt(cipherText))}>DECRYPT</button>
            <button onClick={() => WebassemblyInjection(decryptText)}>INJECTION</button>
            <p>{cipherText}</p>
            <p>{decryptText}</p>
            <div id="testDiv"></div>
        </>
    )
}

export default App;