use reformata_cms_core::app;

#[tokio::main]
async fn main() {
    
    let app = app();
    let listener = tokio::net::TcpListener::bind("0.0.0.0:808000").await.unwrap();

    axum::serve(listener, app).await.unwrap();

}
