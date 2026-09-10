use axum::{
    routing::{get, post},
    Json, Router,
};
use serde::Deserialize;
use tokio::net::TcpListener;

#[derive(Deserialize, Debug)]
struct Report {
    country: String,
    warplanes_in_air: i32,
    warships_in_water: i32,
    timestamp: String,
}

// Ruta necesaria para el Health Check de Google Cloud Load Balancer
async fn health_check() -> &'static str {
    "OK"
}

// Ruta principal que procesa las solicitudes POST de Locust
// Recibe el cuerpo JSON y lo deserializa automaticamente a la estructura Report
async fn handle_report(Json(payload): Json<Report>) -> String {
    // Genera carga intensiva de CPU para superar el umbral del 30% y activar el HPA
    let mut _dummy: u64 = 0;
    for i in 0..5_000_000 {
        // wrapping_add evita el panico por desbordamiento de entero en tipo u64
        _dummy = _dummy.wrapping_add(i);
    }
    
    println!("Recibido reporte de: {}", payload.country);
    format!("Reporte de {} recibido en GKE", payload.country)
}

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/", get(health_check))
        .route("/grpc-2026", post(handle_report));

    let listener = TcpListener::bind("0.0.0.0:8080").await.unwrap();
    println!("API Rust en puerto 8080...");
    axum::serve(listener, app).await.unwrap();
}