<template>
  <footer>
    <div class="container">
      <div class="row footer-info">
        <div class="col footer-item">
          <div class="logo">
            <span class="text-primary">TECH </span>
            <span class="text-secondary">SHOP</span>
          </div>
          <p>800-2345-6789</p>
          <p>support@domain.com</p>
        </div>
        <div class="col footer-item">
          <p class="title">Shop online</p>
          <ul>
            <li>Shop 1</li>
            <li>Shop 2</li>
            <li>Shop 3</li>
            <li>Shop 4</li>
            <li>Shop 5</li>
          </ul>
        </div>
        <div class="col footer-item">
          <p class="title">Customer care</p>
          <ul>
            <li>Contact us</li>
            <li>Return Policy</li>
            <li>FAQs</li>
          </ul>
        </div>
        <div class="col footer-item">
          <p class="title">Subscribe</p>
          <p class="sub-title">
            Sign up to get the latest on sales, new releases and more ...
          </p>
          <div class="subscribe">
            <input type="text" placeholder="Your email" />
            <button>ĐĂNG KÝ</button>
          </div>
        </div>
      </div>
      <div class="row down">
        <div>© 2021 template-homedecor.</div>
        <div class="credit">
          <img
            alt=""
            class="payment-icon"
            data-src="https://img.btdmp.com/files/10088580/2020/06/15/15922126940058f3e369.png"
            src="https://img.btdmp.com/files/10088580/2020/06/15/15922126940058f3e369.png"
            lazy="loaded"
          /><img
            alt=""
            class="ml8 payment-icon"
            data-src="https://img.btdmp.com/files/10088580/2020/06/15/1592212702159c1d36ae.png"
            src="https://img.btdmp.com/files/10088580/2020/06/15/1592212702159c1d36ae.png"
            lazy="loaded"
          /><img
            alt=""
            class="ml8 payment-icon"
            data-src="https://img.btdmp.com/files/10088580/2020/06/15/1592212709e5152d4744.png"
            src="https://img.btdmp.com/files/10088580/2020/06/15/1592212709e5152d4744.png"
            lazy="loaded"
          /><img
            alt=""
            class="ml8 payment-icon"
            data-src="https://img.btdmp.com/files/10088580/2020/06/15/1592212715afd34ddc0e.png"
            src="https://img.btdmp.com/files/10088580/2020/06/15/1592212715afd34ddc0e.png"
            lazy="loaded"
          />
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: "FooterPage",
  props: {},
};
</script>

<style scoped lang="scss">
footer {
  background-color: #fcb800;
  padding: 50px 0;
  .footer-info {
    display: flex;
    padding-bottom: 30px;
    border-bottom: 1px solid rgb(121, 121, 121);
    .footer-item {
      flex: 25%;
      p {
        margin: 0;
      }
      .logo {
        font-size: 20px;
        font-weight: 700;
        padding-bottom: 20px;
      }
      .title {
        font-weight: 700;
      }
      .sub-title {
        padding: 20px 0;
      }
      .subscribe {
        width: 100%;
        display: flex;
        input {
          padding: 10px 5px;
          outline: none;
          border: none;
          flex: 70%;
        }
        button {
          flex: 30%;
          padding: 10px 5px;
          outline: none;
          border: none;
          cursor: pointer;
          transition: all 0.3s ease-in;
          &:hover {
            background-color: rgb(122, 122, 255);
            color: white;
          }
        }
      }
    }
  }
  .down {
    padding-top: 30px;
    display: flex;
    justify-content: space-between;
    .credit {
      img {
        padding-left: 10px;
      }
    }
  }
}
@media (max-width: 768px) {
  footer {
    .footer-info {
      flex-direction: column;
      .footer-item {
        padding-bottom: 30px;
      }
    }
    .down {
      flex-direction: column;
      justify-content: center;
      div {
        text-align: center;
        padding-bottom: 20px;
      }
      .credit {
        img {
          padding-left: 5px;
        }
      }
    }
  }
}
@media (max-width: 576px) {
  footer {
    .container {
      padding: 10px;
    }
  }
}
</style>
