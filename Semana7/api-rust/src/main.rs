use actix_web::{web, App, HttpServer, HttpResponse, Result, middleware::Logger};
use serde::{Deserialize, Serialize};
use std::env;

// Estructura para el json /clima
#[derive(Debug, Serialize, Deserialize)]
struct DatosClima {
    ciudad: String,
    temperatura: f32,
    humedad: f32,
    condicion: String,
}

//handler
async fn hola() -> Result<HttpResponse> {
    log::info!("Endpoint received");
    Ok(HttpResponse::Ok().body("Hi, welcome to Climate API"))
}

// handler para el endpoint /clima
async fn post_clima(datos: web::Json<DatosClima>) -> Result<HttpResponse> {
    log::info!("Received clima data: {:?}", datos);
    
    let oracion = format!(
        "El clima en {} es de {}°C con una humedad del {}% y condición de {}.",
        datos.ciudad, datos.temperatura, datos.humedad, datos.condicion
    );

    log::info!("Generated sentence: {}", oracion);

    Ok(HttpResponse::Ok().json(serde_json::json!({
        "mensaje": oracion,
        "datos_recibidos": datos.into_inner()
    })))

}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::init_from_env(env_logger::Env::new().default_filter_or("info"));

    let port = 8080;
    log::info!("Starting server on port {}", port);

    HttpServer::new(|| {
        App::new()
            .wrap(Logger::default())
            .wrap(Logger::new("%a %{User-Agent}i"))
            .route("/", web::get().to(hola))
            .route("/clima", web::post().to(post_clima))
    })
    .bind(("127.0.0.1", port))?
    .run()
    .await
}