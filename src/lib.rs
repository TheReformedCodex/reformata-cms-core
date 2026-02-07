pub mod routes;
pub mod templates;
pub mod state;

use axum::Router;
use tera::Tera;
use lazy_static::lazy_static;

use crate::state::AppState;


lazy_static! {
        pub static ref TEMPLATES: Tera = {
            let mut tera = match Tera::new("html/**/*.html") {
                Ok(t) => t,
                Err(e) => {
                    println!("Parsing errors: {}", e);
                    ::std::process::exit(1);
                }
            };
            tera.autoescape_on(vec![".html", ".sql"]);
            // tera.register_filter("do_nothing", do_nothing_filter);
            tera
        };
}

pub fn app() -> Router<AppState>{
    Router::new().merge(routes::router())
}
