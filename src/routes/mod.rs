use axum::Router;

use crate::state::AppState;

mod home;


pub fn router() -> Router<AppState> {
    Router::new()
        .merge(home::router())
}