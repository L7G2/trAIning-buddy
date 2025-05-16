import { NavLink, Outlet } from "react-router-dom";
import classes from "./NotLoggedNav.module.css";

function NotLoggedNav() {
  return (
    <>
      <nav className={classes.navigation}>
        <NavLink to="/" className={classes.image}>
          trAIningbuddy
        </NavLink>
        <div className={classes.links}>
          <ul className={classes.ul}>
            <li className={classes.li}>
              <NavLink
                to="about"
                className={({ isActive }) =>
                  isActive ? classes.active : classes.notactive
                }
              >
                O nas
              </NavLink>
            </li>
            <li className={classes.li}>
              <NavLink
                to="coaches"
                className={({ isActive }) =>
                  isActive ? classes.active : classes.notactive
                }
              >
                Nasi trenerzy
              </NavLink>
            </li>
            <li className={classes.li}>
              <NavLink
                to="caloriecalculator"
                className={({ isActive }) =>
                  isActive ? classes.active : classes.notactive
                }
              >
                Kalkulator Kalorii
              </NavLink>
            </li>
            <li className={classes.li}>
              <NavLink
                to="login"
                className={({ isActive }) =>
                  isActive ? classes.active : classes.notactive
                }
              >
                Zaloguj się
              </NavLink>
            </li>
          </ul>
        </div>
      </nav>
      <Outlet />
    </>
  );
}

export default NotLoggedNav;
