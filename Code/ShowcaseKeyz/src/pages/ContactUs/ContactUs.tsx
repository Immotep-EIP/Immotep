import { useTranslation } from "react-i18next";
import { useEffect, useRef, useState } from "react";
import style from "./ContactUs.module.css";

function ContactUsPage() {
  const { t } = useTranslation();
  const titleRef = useRef<HTMLHeadingElement>(null);
  const [isVisible, setIsVisible] = useState(false);

  const [formData, setFormData] = useState({
    firstname: "",
    lastname: "",
    email: "",
    message: "",
  });

  const [formTouched, setFormTouched] = useState(false);

  const [errors, setErrors] = useState({
    firstname: false,
    lastname: false,
    email: false,
    message: false,
  });

  const isFormValid = () => {
    return (
      formData.firstname.trim() !== "" &&
      formData.lastname.trim() !== "" &&
      formData.email.trim() !== "" &&
      formData.message.trim() !== ""
    );
  };

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;

    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));

    setFormTouched(true);

    setErrors((prev) => ({
      ...prev,
      [name]: value.trim() === "",
    }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    const newErrors = {
      firstname: formData.firstname.trim() === "",
      lastname: formData.lastname.trim() === "",
      email: formData.email.trim() === "",
      message: formData.message.trim() === "",
    };

    setErrors(newErrors);

    if (isFormValid()) {
      console.log("Form data submitted:", formData);
      setFormData({
        firstname: "",
        lastname: "",
        email: "",
        message: "",
      });
      setFormTouched(false);
      alert("Votre message a été envoyé avec succès!");
    } else {
      console.log("Form validation failed");
    }
  };

  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setIsVisible(true);
        } else {
          setIsVisible(false);
        }
      },
      {
        threshold: 0.3,
        rootMargin: "0px 0px -100px 0px",
      }
    );
    if (titleRef.current) {
      observer.observe(titleRef.current);
    }
    return () => {
      if (titleRef.current) {
        observer.unobserve(titleRef.current);
      }
    };
  }, []);

  return (
    <div className={style.pageContainer}>
      <div className={style.titleContainer}>
        <h1 ref={titleRef} className={style.title}>
          {t("contact_us.title")}
        </h1>
        <div
          className={`${style.titleUnderline} ${
            isVisible ? style.animate : ""
          }`}
        />
      </div>

      <div className={style.contentContainer}>
        <div className={style.formContainer}>
          <form className={style.form} onSubmit={handleSubmit}>
            <div className={style.rowContainer}>
              <div className={style.inputGroup}>
                <label htmlFor="firstname" className={style.label}>
                  {t("contact_us.firstname")}
                </label>
                <input
                  type="text"
                  id="firstname"
                  name="firstname"
                  value={formData.firstname}
                  onChange={handleChange}
                  className={`${style.input} ${
                    errors.firstname ? style.inputError : ""
                  }`}
                  required
                />
                {errors.firstname && formTouched && (
                  <span className={style.errorMessage}>
                    {t("contact_us.field_required")}
                  </span>
                )}
              </div>
              <div className={style.inputGroup}>
                <label htmlFor="lastname" className={style.label}>
                  {t("contact_us.lastname")}
                </label>
                <input
                  type="text"
                  id="lastname"
                  name="lastname"
                  value={formData.lastname}
                  onChange={handleChange}
                  className={`${style.input} ${
                    errors.lastname ? style.inputError : ""
                  }`}
                  required
                />
                {errors.lastname && formTouched && (
                  <span className={style.errorMessage}>
                    {t("contact_us.field_required")}
                  </span>
                )}
              </div>
            </div>
            <div className={style.inputGroup}>
              <label htmlFor="email" className={style.label}>
                {t("contact_us.your_email")}
              </label>
              <input
                type="email"
                id="email"
                name="email"
                value={formData.email}
                onChange={handleChange}
                className={`${style.input} ${
                  errors.email ? style.inputError : ""
                }`}
                required
              />
              {errors.email && formTouched && (
                <span className={style.errorMessage}>
                  {t("contact_us.field_required")}
                </span>
              )}
            </div>
            <div className={style.inputGroup}>
              <label htmlFor="message" className={style.label}>
                {t("contact_us.message")}
              </label>
              <textarea
                id="message"
                name="message"
                value={formData.message}
                onChange={handleChange}
                className={`${style.textarea} ${
                  errors.message ? style.inputError : ""
                }`}
                required
              ></textarea>
              {errors.message && formTouched && (
                <span className={style.errorMessage}>
                  {t("contact_us.field_required")}
                </span>
              )}
            </div>
            <button
              type="submit"
              className={`${style.submitButton} ${
                !isFormValid() || !formTouched ? style.submitDisabled : ""
              }`}
              disabled={!isFormValid() || !formTouched}
            >
              <span>{t("contact_us.send_message")}</span>
              <svg className={style.sendIcon} viewBox="0 0 24 24" fill="none">
                <path
                  d="M22 2L11 13"
                  stroke="currentColor"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
                <path
                  d="M22 2l-7 20-4-9-9-4 20-7z"
                  stroke="currentColor"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
            </button>
          </form>
        </div>
      </div>
      <div className={style.decorWave}></div>
      <div className={style.decorCircle1}></div>
      <div className={style.decorCircle2}></div>
      <div className={style.decorDots}></div>
    </div>
  );
}

export default ContactUsPage;
