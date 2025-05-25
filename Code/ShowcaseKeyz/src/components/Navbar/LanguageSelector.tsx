import { useState } from "react";
import { useTranslation } from "react-i18next";
import { ChevronDown } from "lucide-react";
import style from "./LanguageSelector.module.css";

const LanguageSelector = () => {
  const { i18n } = useTranslation();
  const [isOpen, setIsOpen] = useState(false);
  const [selectedLanguage, setSelectedLanguage] = useState("🇫🇷 FR");

  const changeLanguage = (lng: string, label: string) => {
    i18n.changeLanguage(lng);
    setSelectedLanguage(label);
    setIsOpen(false);
  };

  return (
    <div className={style.languageSelector}>
      <button onClick={() => setIsOpen(!isOpen)} className={style.button}>
        {selectedLanguage}
        <ChevronDown className={style.icon} />
      </button>

      {isOpen && (
        <div className={style.dropdown}>
          <ul className={style.list}>
            <li>
              <button
                onClick={() => changeLanguage("fr", "🇫🇷 FR")}
                className={style.listItem}
              >
                🇫🇷 FR
              </button>
            </li>
            <li>
              <button
                onClick={() => changeLanguage("en", "🇺🇸 EN")}
                className={style.listItem}
              >
                🇺🇸 EN
              </button>
            </li>
            <li>
              <button
                onClick={() => changeLanguage("de", "🇩🇪 DE")}
                className={style.listItem}
              >
                🇩🇪 DE
              </button>
            </li>
          </ul>
        </div>
      )}
    </div>
  );
};

export default LanguageSelector;
