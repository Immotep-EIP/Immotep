import { useTranslation } from "react-i18next";
import { useEffect, useRef, useState } from "react";
import style from "./ContactUs.module.css";

function ContactUsPage() {
  const { t } = useTranslation();
  const titleRef = useRef<HTMLHeadingElement>(null);
  const [isVisible, setIsVisible] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [notification, setNotification] = useState<{
    show: boolean;
    type: "success" | "error";
    message: string;
  }>({
    show: false,
    type: "success",
    message: "",
  });

  const [formData, setFormData] = useState({
    firstname: "",
    lastname: "",
    email: "",
    subject: "",
    message: "",
  });

  const [formTouched, setFormTouched] = useState(false);

  const [errors, setErrors] = useState({
    firstname: false,
    lastname: false,
    email: false,
    subject: false,
    message: false,
  });

  const isFormValid = () => {
    return (
      formData.firstname.trim() !== "" &&
      formData.lastname.trim() !== "" &&
      formData.email.trim() !== "" &&
      formData.subject.trim() !== "" &&
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

  const showNotification = (type: "success" | "error", message: string) => {
    setNotification({
      show: true,
      type,
      message,
    });

    setTimeout(() => {
      setNotification((prev) => ({ ...prev, show: false }));
    }, 5000);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const newErrors = {
      firstname: formData.firstname.trim() === "",
      lastname: formData.lastname.trim() === "",
      email: formData.email.trim() === "",
      subject: formData.subject.trim() === "",
      message: formData.message.trim() === "",
    };

    setErrors(newErrors);

    if (isFormValid()) {
      setIsLoading(true);

      try {
        const response = await fetch(import.meta.env.VITE_API_URL + "/contact/", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(formData),
        });

        if (!response.ok) {
          throw new Error("Failed to send message");
        }

        const result = await response.json();
        console.log("API response:", result);

        setFormData({
          firstname: "",
          lastname: "",
          email: "",
          subject: "",
          message: "",
        });
        setFormTouched(false);
        showNotification("success", t("contact_us.success_message"));
      } catch (error) {
        console.error("Error submitting form:", error);
        showNotification("error", t("contact_us.error_message"));
      } finally {
        setIsLoading(false);
      }
    } else {
      console.log("Form validation failed");
      showNotification("error", t("contact_us.validation_error"));
    }
  };

  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        setIsVisible(entry.isIntersecting);
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

      {notification.show && (
        <div className={`${style.notification} ${style[notification.type]}`}>
          <div className={style.notificationContent}>
            {notification.type === "success" ? (
              <svg
                className={style.notificationIcon}
                viewBox="0 0 24 24"
                fill="none"
              >
                <path
                  d="M22 11.08V12a10 10 0 1 1-5.93-9.14"
                  stroke="currentColor"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
                <path
                  d="M22 4L12 14.01l-3-3"
                  stroke="currentColor"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
            ) : (
              <svg
                className={style.notificationIcon}
                viewBox="0 0 24 24"
                fill="none"
              >
                <circle
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  strokeWidth="2"
                />
                <path
                  d="M12 8v4M12 16h.01"
                  stroke="currentColor"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
            )}
            <span>{notification.message}</span>
            <button
              className={style.notificationClose}
              onClick={() =>
                setNotification((prev) => ({ ...prev, show: false }))
              }
            >
              <svg viewBox="0 0 24 24" fill="none">
                <path
                  d="M18 6L6 18M6 6l12 12"
                  stroke="currentColor"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
            </button>
          </div>
        </div>
      )}

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
                  placeholder={t("contact_us.firstname_placeholder")}
                  disabled={isLoading}
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
                  placeholder={t("contact_us.lastname_placeholder")}
                  disabled={isLoading}
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
                placeholder={t("contact_us.your_email_placeholder")}
                disabled={isLoading}
              />
              {errors.email && formTouched && (
                <span className={style.errorMessage}>
                  {t("contact_us.field_required")}
                </span>
              )}
            </div>
            <div className={style.inputGroup}>
              <label htmlFor="subject" className={style.label}>
                {t("contact_us.subject")}
              </label>
              <input
                type="text"
                id="subject"
                name="subject"
                value={formData.subject}
                onChange={handleChange}
                className={`${style.input} ${
                  errors.subject ? style.inputError : ""
                }`}
                required
                placeholder={t("contact_us.subject_placeholder")}
                disabled={isLoading}
              />
              {errors.subject && formTouched && (
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
                placeholder={t("contact_us.message_placeholder")}
                disabled={isLoading}
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
                !isFormValid() || !formTouched || isLoading
                  ? style.submitDisabled
                  : ""
              } ${isLoading ? style.loading : ""}`}
              disabled={!isFormValid() || !formTouched || isLoading}
            >
              {isLoading ? (
                <div className={style.spinner}></div>
              ) : (
                <>
                  <span>{t("contact_us.send_message")}</span>
                  <svg
                    className={style.sendIcon}
                    viewBox="0 0 24 24"
                    fill="none"
                  >
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
                </>
              )}
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
