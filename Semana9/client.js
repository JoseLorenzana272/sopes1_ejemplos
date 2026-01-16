const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const path = require('path');

// cargar el archivo .proto
const packageDefinition = protoLoader.loadSync(
    path.resolve(__dirname, 'proto/sysinfo.proto'),
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    }
);

// cargar paquete sysinfo
const sysinfoProto = grpc.loadPackageDefinition(packageDefinition).sysinfo;

// crear ciente
const client = new sysinfoProto.SystemMonitor(
    'localhost:50051',
    grpc.credentials.createInsecure()
)

// llamada
console.log("[Node Client] Conectando al servidor Go");

client.GetRamStatus({ host_name: "Cliente-NodeJS" }, (err, response) => {
    if (err) {
        console.error("Error:", err);
    } else {
        console.log("\n Respuesta del Servidor:");
        console.log(`   Mensaje: ${response.message}`);
        console.log(`   RAM Usada: ${response.used_percent}%`);
    }
});
