<template>
  <main id="main">
    <h1 id="title">Edward Jenner</h1>
    <p>The scientist who pioneered the concept of vaccines</p>
    <!-- Image of Edward Jenner -->
    <figure id="img-div">
      <img
        id="image"
        src="https://upload.wikimedia.org/wikipedia/commons/0/0c/Edward_Jenner._Oil_painting._Wellcome_V0023503.jpg"
        alt="Dr. Norman Borlaug seen standing in Mexican wheat field with a group of biologists"
      />
      <figcaption id="img-caption">Oilpainting of Edward Jenner.</figcaption>
    </figure>
    <!-- End image of Edward Jenner -->
    <!-- Content section -->
    <section id="tribute-info">
      <h2>Invention of the vaccine</h2>
      <p>
        Inoculation (which is the process of producing immunity by introduction
        of the infectious agent onto an abraded or absorptive skin surface
        instead of inserting the substance in the tissues by injection.) was
        already pioneered in Asian medicine and was a standard practice but
        involved serious risks, one of which was the fear that those inoculated
        would then transfer the disease to those around them due to them
        becoming carriers of the disease. In 1721, Lady Mary Wortley Montagu had
        imported variolation to Britain after having observed it in
        Constantinople. While Johnnie Notions had great success with his
        self-devised inoculation, his method's practice was limited to the
        Shetland Isles. Voltaire wrote that at this time 60% of the population
        caught smallpox and 20% of the population died of it. Voltaire also
        states that the Circassians used the inoculation from times immemorial,
        and the custom may have been borrowed by the Turks from the Circassians.
        In 1766, Daniel Bernoulli analysed smallpox morbidity and mortality data
        to demonstrate the efficacy of inoculation.
      </p>

      <p>
        By 1768, English physician John Fewster had realised that prior
        infection with cowpox rendered a person immune to smallpox. In the years
        following 1770, at least five investigators in England and Germany
        (Sevel, Jensen, Jesty 1774, Rendell, Plett 1791) successfully tested in
        humans a cowpox vaccine against smallpox. For example, Dorset farmer
        Benjamin Jesty successfully vaccinated and presumably induced immunity
        with cowpox in his wife and two children during a smallpox epidemic in
        1774, but it was not until Jenner's work that the procedure became
        widely understood. Jenner may have been aware of Jesty's procedures and
        success. A similar observation was later made in France by Jacques
        Antoine Rabaut-Pommier in 1780.
      </p>

      <p>
        Noting the common observation that milkmaids were generally immune to
        smallpox, Jenner postulated that the pus in the blisters that milkmaids
        received from cowpox (a disease similar to smallpox, but much less
        virulent) protected them from smallpox.
      </p>

      <p>
        On 14 May 1796, Jenner tested his hypothesis by inoculating James
        Phipps, an eight-year-old boy who was the son of Jenner's gardener. He
        scraped pus from cowpox blisters on the hands of Sarah Nelmes, a
        milkmaid who had caught cowpox from a cow called Blossom. Phipps was the
        17th case described in Jenner's first paper on vaccination.
      </p>

      <p>
        Jenner inoculated Phipps in both arms that day, subsequently producing
        in Phipps a fever and some uneasiness, but no full-blown infection.
        Later, he injected Phipps with variolous material, the routine method of
        immunization at that time. No disease followed. The boy was later
        challenged with variolous material and again showed no sign of
        infection.
      </p>

      <p>
        Donald Hopkins has written, "Jenner's unique contribution was not that
        he inoculated a few persons with cowpox, but that he then proved [by
        subsequent challenges] that they were immune to smallpox. Moreover, he
        demonstrated that the protective cowpox pus could be effectively
        inoculated from person to person, not just directly from cattle." Jenner
        successfully tested his hypothesis on 23 additional subjects.
      </p>

      <p>
        Jenner continued his research and reported it to the Royal Society,
        which did not publish the initial paper. After revisions and further
        investigations, he published his findings on the 23 cases, including his
        11-month-old son Robert. Some of his conclusions were correct, some
        erroneous. Modern microbiological and microscopic methods would make his
        studies easier to reproduce. The medical establishment deliberated at
        length over his findings before accepting them. Eventually, vaccination
        was accepted, and in 1840, the British government banned variolation –
        the use of smallpox to induce immunity – and provided vaccination using
        cowpox free of charge.
      </p>

      <p>
        The success of his discovery soon spread around Europe and was used en
        masse in the Spanish Balmis Expedition (1803–1806), a three-year-long
        mission to the Americas, the Philippines, Macao, China, led by Dr.
        Francisco Javier de Balmis with the aim of giving thousands the smallpox
        vaccine. The expedition was successful. Napoleon, who at the time was at
        war with Britain, had all his French troops vaccinated, awarded Jenner a
        medal, and at the request of Jenner, he released two English prisoners
        of war and permitted their return home.
      </p>

      <p>
        Jenner's continuing work on vaccination prevented him from continuing
        his ordinary medical practice. He was supported by his colleagues and
        the King in petitioning Parliament, and was granted £10,000 in 1802 for
        his work on vaccination. In 1807, he was granted another £20,000 after
        the Royal College of Physicians confirmed the widespread efficacy of
        vaccination.
      </p>

      <blockquote
        cite="https://www.brainyquote.com/quotes/edward_jenner_190635"
      >
        <p>
          I hope that some day the practice of producing cowpox in human beings
          will spread over the world - when that day comes, there will be no
          more smallpox.
        </p>
        <cite>- Edward Jenner</cite>
      </blockquote>
    </section>
    <!-- End content section-->
    <!-- Reference section-->
    <section class="references">
      <h2>References</h2>
      <a
        class="tribute-link"
        id="tribute-link"
        href="https://en.wikipedia.org/wiki/Edward_Jenner"
        target="_blank"
        >Wikipedia</a
      >

      <a
        class="tribute-link"
        href="https://www.britannica.com/science/inoculation"
        target="_blank"
        >Britannica</a
      >
    </section>
    <!-- End reference section-->
  </main>
</template>

<script>
export default {};
</script>

<style scoped>
body {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto",
    "Helvetica Neue", Arial, sans-serif;
  font-size: 1.7rem;
  line-height: 1.5;
  text-align: center;
  color: #333;
  margin: auto;
}

section {
  margin-top: 10px;
  margin-bottom: 30px;
}

.references {
  background: #ffc508;
  padding-top: 10px;
  padding-bottom: 30px;
}

h1 {
  font-size: 4rem;
  margin-bottom: 0;
}

@media (max-width: 460px) {
  h1 {
    font-size: 3.5rem;
    line-height: 1.2;
  }
}

h2 {
  margin: 25px 0;
  text-align: center;
  font-size: 2.5rem;
}

#tribute-info p {
  margin-left: 20%;
  margin-right: 20%;
  text-align: left;
}

a:link {
  color: black;
}

a:visited {
  color: brown;
}

a:hover {
  color: red;
}

a:active {
  color: white;
}

img {
  max-width: 25%;
  display: block;
  border-radius: 50%;
  height: auto;
  margin: 0 auto;
}

#img-div {
  background: rgb(255, 133, 66);
  background: linear-gradient(
    0deg,
    rgba(255, 133, 66, 1) 0%,
    rgba(255, 198, 55, 1) 100%
  );
  padding: 20px;
  margin: 0;
}

#img-caption {
  margin: 20px 0 5px 0;
  font-size: 1.8rem;
}

@media (max-width: 460px) {
  #img-caption {
    font-size: 1.4rem;
  }
}

blockquote {
  font-style: italic;
  max-width: 545px;
  margin: 0 auto 50px auto;
  text-align: left;
}
</style>
