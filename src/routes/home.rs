use axum::{
    extract::State,
    routing::get, 
    Router,
    response::Html
};
use tera::Context;
use crate::state::{self, AppState};

// use crate


// pub async fn home() -> IndexTemplate

pub fn router() -> Router<AppState> {
    Router::new().route("/", get(index))
}


async fn index(State(state): State<AppState>) -> Html<String> {

    let mut ctx = Context::new();

    ctx.insert("title", "Home");
    ctx.insert("message", "Hello from axum and tera");

    let rendered = state
        .tera.render("home/home.html", &ctx)
        .unwrap();

    Html(rendered)
}